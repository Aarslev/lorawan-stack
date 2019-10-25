// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import axios from 'axios'
import TTN from 'ttn-lw'

import token from '../lib/access-token'
import getCookieValue from '../../lib/cookie'
import { selectApplicationConfig, selectApplicationRootPath } from '../../lib/selectors/env'

const config = selectApplicationConfig()
const appRoot = selectApplicationRootPath()

const stack = {
  is: config.is.enabled ? config.is.base_url : undefined,
}

const isBaseUrl = config.is.base_url

const ttnClient = new TTN(token, {
  stackConfig: stack,
  connectionType: 'http',
  proxy: false,
})

const csrf = getCookieValue('_csrf')
const instance = axios.create({
  headers: { 'X-CSRF-Token': csrf },
})

export default {
  claim: {
    token() {
      return instance.get(`${appRoot}/api/auth/token`)
    },
    logout() {
      return instance.post(`${appRoot}/api/auth/logout`)
    },
  },
  clients: {
    get(client_id) {
      return instance.get(`${isBaseUrl}/is/clients/${client_id}`)
    },
  },
  users: {
    async get(userId) {
      return instance.get(`${isBaseUrl}/users/${userId}`, {
        headers: {
          Authorization: `Bearer ${(await token()).access_token}`,
        },
      })
    },
    async authInfo() {
      return instance.get(`${isBaseUrl}/auth_info`, {
        headers: {
          Authorization: `Bearer ${(await token()).access_token}`,
        },
      })
    },
  },
  applications: {
    list: ttnClient.Applications.getAll.bind(ttnClient.Applications),
    search: ttnClient.Applications.search.bind(ttnClient.Applications),
  },
  application: {
    get: ttnClient.Applications.getById.bind(ttnClient.Applications),
    delete: ttnClient.Applications.deleteById.bind(ttnClient.Applications),
    create: ttnClient.Applications.create.bind(ttnClient.Applications),
    update: ttnClient.Applications.updateById.bind(ttnClient.Applications),
    eventsSubscribe: ttnClient.Applications.openStream.bind(ttnClient.Applications),
    apiKeys: {
      get: ttnClient.Applications.ApiKeys.getById.bind(ttnClient.Applications.ApiKeys),
      list: ttnClient.Applications.ApiKeys.getAll.bind(ttnClient.Applications.ApiKeys),
      update: ttnClient.Applications.ApiKeys.updateById.bind(ttnClient.Applications.ApiKeys),
      delete: ttnClient.Applications.ApiKeys.deleteById.bind(ttnClient.Applications.ApiKeys),
      create: ttnClient.Applications.ApiKeys.create.bind(ttnClient.Applications.ApiKeys),
    },
  },
}
