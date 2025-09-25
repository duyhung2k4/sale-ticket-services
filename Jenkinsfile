pipeline{
    agent {
        label "172.17.8.248"
    }
    stages{
        stage("build-iam"){
            steps{
                echo "start build image iam"
                script {
                    sh """#!/bin/bash -l
                        cd /root/project/sale-ticket/sale-ticket-service/iam-service/
                        docker build -f Dockerfile -t iam .
                        docker tag iam nguyen040904/iam
                        docker push nguyen040904/iam
                    """
                }
                echo "end build image iam"
            }
        }
        stage("deploy-iam"){
            steps{
                echo "start deploy iam"
                script {
                    sh """
                        cd /root/project/sale-ticket/sale-ticket-service/k8s/iam/
                        zx service.mjs
                    """
                }
            }
        }
    }
    post{
        always{
            echo "========always========"
        }
        success{
            echo "========pipeline executed successfully ========"
        }
        failure{
            echo "========pipeline execution failed========"
        }
    }
}