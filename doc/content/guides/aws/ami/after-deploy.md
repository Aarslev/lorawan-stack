---
title: "AWS Marketplace AMI"
description: ""
weight: 40
---

This section provides the information necessary to get started with The Things Enterprise Stack after it is deployed via AWS Marketplace using AWS CloudFormation.

## DNS configuration

In order to access the Console/API of the The Things Enterprise Stack, the domain name chosen during deployment must be mapped to the public IP address of the CloudFormation stack.

Login to the [AWS Console](https://console.aws.amazon.com/) and open the **CloudFormation** resource. Click on the CloudFormation stack that you deployed and navigate to the **Outputs** tab and copy the value of the **PublicIP** field.

Map the **Domain** value that was chosen during the deployment to this IP address using the name server that manages this domain. This is typically accomplished by setting an **A** record as follows:

|Name|Record Type|Value|
|---|---|---|
|@ (or sub-domain value)|`A`|**PublicIP**|

This DNS mapping takes some time to propagate. You can cross check this by using `nslookup` on the terminal or a similar domain lookup service online.

Once it is confirmed that the Domain resolves to the **PublicIP**, the Console of The Things Enterprise Stack can be accessed at `https://<domain>/console`.

### SSL certificates

As a security measure, plaintext access to the Console/API is disabled. In order to serve requests securely, The Things Enterprise Stack has built-in support to automatically request, serve and renew SSL certificates.

Apart from the DNS mapping above, no additional steps are necessary for this.

## Getting started using the Console

Please check the [Console Guide]({{< relref "../../getting-started/console/_index.md" >}}) to get quickly started with The Things Enterprise Stack Console.

## Getting started using the Command Line Interface (CLI)

Please check the [CLI Guide]({{< relref "../../getting-started/cli/_index.md" >}}) to get quickly started with The Things Enterprise Stack CLI.

## SSH access

In most cases, you will interact with The Things Enterprise Stack only via the Console or the CLI. However, in case there is a need to directly access the EC2 instance, there is an option to do so via SSH.

SSH access is possible only via the IP addresses set using the **Restrict SSH Access to IP Range** parameter during deployment. Also, the SSH client needs access to the private key corresponding to the public key that was chosen in the **SSH Key** parameter during deployment. The SSH username is `ec2-user`.

For example, you can use the OpenSSH client via the terminal and login using:

```bash
$ ssh -i <private-key-file> ec2-user@PublicIP
```

Upon accessing the machine, navigate to the `tti` directory:

```bash
$ cd /tti
```

This directory is structured as follows:

```
├── acme                        # contains SSL certificates and related files
├── bootstrap                   # contains logs generated during deployment and restarts
└── lorawan-stack
    ├── config.yml              # The Things Enterprise Stack configuration
    ├── environment             # additional environment variables
    ├── lorawan-frequency-plans # contains a local copy of the LoRaWAN Frequency Plans Repository
    ├── public                  # contains the static assets
    └── tti-lw-stack            # The Things Enterprise Stack binary
```

The Things Enterprise Stack binary is run as a `systemd` service. In order to check the logs, run the following:

```bash
$ sudo journalctl -f -u lorawan-stack.service
```

## Updating the CloudFormation stack

We recommend using [Change Sets](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-changesets.html) to to update the CloudFormation Stack.

On the AWS Console, open the **CloudFormation** service, navigate to the **Change sets** tab and select **Create change set**. There are two choices that can be made here:

- Use current template: Choose this option to update input parameters on an existing template.
- Replace current template: Choose this option to use a new template on the existing deployment. This option should be used to apply new versions of The Things Enterprise Stack that are delivered as new version of the CloudFormation template.

Once the necessary options are updated, select the **Create Change Set** option. This will create a change set that describes the resources that will be updated by this change and if any of these resources need to be replaced. 

After confirming the changes, select **Execute**. Depending on the resources, this will take some time to complete and can be tracked using the **Events** tab on the AWS Console.

## Manual upgrade of databases (Amazon RDS/ Redis) using snapshots

The update of certain fields of the CloudFormation stack necessitates the recreation of the database. AWS does not automatically migrate the data and hence, this must be done manually. A disclaimer is added to such fields in the CloudFormation template description.

In order to migrate the database without the loss of data, first create a snapshot of your database before running a change set on these fields.

- For RDS, navigate to **RDS** > **Snapshots** on the AWS Console and select **Take Snapshot**.
- For Redis (ElastiCache), navigate to **ElastiCache** on the AWS Console and select your Redis Replication Group and select **Backup** option.

Now run a change set on CloudFormation.

- For RDS, enter the ARN (Amazon Resource Name) of the snapshot into the **Amazon RDS Snapshot** field and run the change set.
- For Redis (ElastiCache), enter the Name of the backup into the **Amazon RDS Database Name** and run the change set.
