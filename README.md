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

# Installation and use
1. Go to the feature/docker_app branch of this repository
2. Check the ```txns.csv``` file in the resource folder, which contains the transactions to be processed
3. Check the ```.env``` file, which contains the destination email for the balance that the application will generate
4. Locate the root folder and run the script
>./run.sh

### AWS Lambda

