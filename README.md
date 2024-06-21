# Stori - Technical Challenge - Transactions
## Technology used
1. Golang 1.2.1
2. gotodoenv 1.5.1
3. gomail 2.0.0
4. AWS lambda 1.47.0 for lambda implementation
5. AWS SDK Go 1.54.6 for lambda implementation
5. Docker engine for container deployment

## Requirements
1. Golang installed on the machine
2. Docker engine installed on the machine

## Challenge Description
For this challenge you must create a system that processes a file from a mounted directory. The file
will contain a list of debit and credit transactions on an account. Your function should process the file
and send summary information to a user in the form of an email.
An example file is shown below; but create your own file for the challenge. Credit transactions are
indicated with a plus sign like +60.5. Debit transactions are indicated by a minus sign like -20.46

![image](https://github.com/oz0726/StoriTxChallenge/assets/88631143/c577d0ab-6f8e-4032-a022-699b5bd32f9f)

The summary email contains information on the total balance in the account, the number of
transactions grouped by month, and the average credit and average debit amounts grouped by
month. Using the transactions in the image above as an example, the summary info would be

Total balance is 39.74
Number of transactions in July: 2
Number of transactions in August: 2
Average debit amount: -15.38
Average credit amount: 35.25

![image](https://github.com/oz0726/StoriTxChallenge/assets/88631143/55b8ae96-b929-4454-a415-9c854fcba576)


## Solution description
First, a shared project structure was generated for the proposed solutions that has the configuration of the input and output connections of the system and the business logic.
Based on this initial structure, two different approaches are generated for the application:

### Docker APP

We have a resources directory where the ```.csv``` file to be processed is stored. We also have a ```.env``` file where the destination email of the balance to be generated is configured. Finally, we have the ```Dockerfile``` that creates the image of our application and a ```run.sh``` executable file that is responsible for compiling the code, generating the docker image, creating the container, starting the application and stopping and eliminating the container as soon as the execution ends.

#### Installation and use
1. Go to the ```feature/docker_app``` branch of this repository
2. Check the ```txns.csv``` file in the resource folder, which contains the transactions to be processed
3. Check the ```.env``` file, which contains the destination email for the balance that the application will generate
4. Locate the root folder and run the script
>./run.sh
5. If the configuration is correct, the script should generate the balance by email.

### AWS Lambda

We have a ```.env``` file where the destination email of the balance to be generated is configured. Also, we have a ```package.sh``` executable that is responsible for compiling the code and generating the appropriate executable to load into the AWS lambda.

#### AWS lambda base configuration
1. Runtime configuration: Amazon Linux 2023
2. Architecture: x86_64
3. Environment variables to configure:
   
```SMTP_SENDER (email from which the email will be generated in the smtp)```

```SMTP_PASSWORD (password to use the smtp service)```

```AWS_BUCKET (container that will have the .csv file to be processed)```

```AWS_OBJECT (name of the .csv file to be processed)```

Additionally, permissions must be given to the role associated with the lambda to be able to read the files in the S3 bucket.
Finally, the triggers must be configured to start the lambda, in the example I created a trigger when the transaction file is updated and another to expose a get endpoint that triggers the generation of the balance with the file that is in the bucket

![image](https://github.com/oz0726/StoriTxChallenge/assets/88631143/aeff579b-410f-403d-85a8-8ebaf616c002)


#### Installation and use
1. Go to the ```feature/aws_lambda``` branch of this repository
2. Check the ```txns.csv``` file in the resource folder, which contains the transactions to be processed
3. Check the ```.env``` file, which contains the destination email for the balance that the application will generate
4. Locate the root folder and run the script

>./package.sh

5. Create a .zip file that has the following files

```Bootstrap``` file located in the root of the project

```.env``` file located in the root of the project

```resources``` folder located in the root of the project

6. In AWS, create a lambda with the indicated configuration
7. In AWS, create an s3 bucket and place the ```txns.csv``` file to be processed there
8. In the code configuration of the created lambda, upload the generated .zip

![image](https://github.com/oz0726/StoriTxChallenge/assets/88631143/42b27035-296c-4dc3-9295-077b87984bb6)

![image](https://github.com/oz0726/StoriTxChallenge/assets/88631143/c18c991e-3bce-4d6c-b4b9-badaa29bbc90)



9. If the configuration is correct, the triggers configured in the lambda should generate the balance by email

Author: [Olman Ibañez](https://www.linkedin.com/in/olman-ibanez/)
