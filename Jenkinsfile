pipeline {

    environment {
        nameImage = "docker.nimtechnology.com/hometest/hometest-golang-tiki"
        privateRegistry = "https://docker.nimtechnology.com"
        registryCredential = "docker_nimtechnology"
        nameImagePublic = "mrnim94/hometest-golang-tiki"
        publicRegistry = "https://registry.hub.docker.com"
        registryCredentialPublic = "public-docker-hub"
    }

    agent any
    
    parameters {
        gitParameter name: 'BRANCH',
            type: 'PT_BRANCH',
            defaultValue: 'master'
    }
    
    stages {
        stage('Check GIT') {
            steps {
                checkout([$class: 'GitSCM',
                    branches: [[name: "${params.BRANCH}"]],
                    doGenerateSubmoduleConfigurations: false,
                    extensions: [],
                    gitTool: 'Default',
                    submoduleCfg: [],
                    userRemoteConfigs: [[credentialsId: 'Jenkin-login-Gitlab', url: 'https://gitlab.nimtechnology.com/nim/home-test-tiki.git']]
                ])
            }
        }

        // stage('Git Clone file environment') {
        //     steps {
        //         withCredentials([[$class: 'UsernamePasswordMultiBinding', credentialsId:'Jenkin-login-Gitlab', usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD']]) {
        //             sh "rm -rf ./env-selenium-check-awingu"
        //             sh "git clone https://$USERNAME:$PASSWORD@gitlab.nimtechnology.com/nim/env-selenium-check-awingu.git"
        //             sh "cp ./env-selenium-check-awingu/config_file/.env ./config_file/.env"
        //             sh "cp ./env-selenium-check-awingu/start-golang.sh ./cmd/production/start-golang.sh"
        //             sh "rm -rf ./env-selenium-check-awingu"
        //         }
        //     }
        // }

        // stage('Code Quality Check via SonarQube') {
        //     steps {
        //        script {
        //        def scannerHome = tool 'sonarqube-scanner';
        //            withSonarQubeEnv("sonarqube-container") {
        //            sh "${tool("sonarqube-scanner")}/bin/sonar-scanner \
        //            -Dsonar.projectKey=selenium-check-awingu \
        //            -Dsonar.sources=. \
        //            -Dsonar.host.url=http://192.168.101.5:9000 \
        //            -Dsonar.login=a4abc687d2f4249e77ea127173c2643a7b8f8300"
        //                }
        //         }
        //     }
        // }
        // thang them
        // test1 co thay khong?

        stage('Prepare ENV for Build') {
            steps {
                script {
                    //IMAGE_NAME = "docker.nimtechnology.com/news-nimtechnology/db"
                    GIT_BRANCH = sh(returnStdout: true, script: "git rev-parse --abbrev-ref HEAD").trim()
                    GIT_HASH = sh(returnStdout: true, script: "git rev-parse --short HEAD").trim()
                }
            // echo "branch is: ${env.GIT_BRANCH}"
            // echo "hash is: ${GIT_HASH}"
            // echo "hash is: ${IMAGE_NAME}"
            sh "echo GIT_HASH=${GIT_HASH} > trigger.properties"
            sh "echo ACTION=DEPLOY >> trigger.properties"
            sh "echo TAG=1.5-SNAPSHOT >> trigger.properties"
            archiveArtifacts 'trigger.properties'
            }
        }

        stage('Create images') {
            steps {
                // sh "docker build -t ${IMAGE_NAME}:latest ."
                // sh "docker build -t ${IMAGE_NAME}:${GIT_HASH} ."
                script {
                  app_latest = docker.build nameImage + ":latest"
                  app_version = docker.build nameImage + ":$GIT_HASH"
                }
            }
        }

        stage('Push Image Private HUB Docker') {
            steps{
                script {
                    docker.withRegistry( privateRegistry, registryCredential ) {
                        app_latest.push()
                        app_version.push()
                    }
                }
            }
        }

        stage('Create images for public') {
            steps {
                // sh "docker build -t ${IMAGE_NAME}:latest ."
                // sh "docker build -t ${IMAGE_NAME}:${GIT_HASH} ."
                script {
                  app_latest_public = docker.build nameImagePublic + ":latest"
                  app_version_public = docker.build nameImagePublic + ":$GIT_HASH"
                }
            }
        }

        stage('Push Image Public HUB Docker') {
            steps{
                script {
                    docker.withRegistry( publicRegistry, registryCredentialPublic ) {
                        app_latest_public.push()
                        app_version_public.push()
                    }
                }
            }
        }

        stage('Remove Unused docker image') {
            steps{
                sh "docker rmi $nameImage:latest"
                sh "docker rmi $nameImage:$GIT_HASH"
                sh "docker rmi $nameImagePublic:latest"
                sh "docker rmi $nameImagePublic:$GIT_HASH"
            }
        }

        stage('Remove Dangling Images') {
            steps{
                sh " docker system prune -f"
            }
        }
    }
}
