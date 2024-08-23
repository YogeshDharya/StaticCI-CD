pipeline{
    agent any 
    environment{
        DOCKERHUB_CREDENTIALS= 
        SECOND_EC2_SSH_CREDENTIALS=
    }
    stages {
        stage("Load Environment Variables"){
            steps{
                script{
                    def envFile = readFile
                }
            }
        }
        stage("Build Docker Image"){
            steps{
                script{
                    docker.build("myusername/myap:${env.BUILD_ID}")
                }
            }
        }
        stage("test"){
            steps{
                echo "At the Testing Stage . This shall succeed !"
            }
        }
        stage("Push Docker Image to DockerHub"){
            steps{
                script{
                    docker.withRegistry('https://registry.hub.docker.com',DOCKERHUB_CREDENTIALS)
                    docker.image("myusername/myapp:${env.BUILD_ID}").push()
                }
            }
        }
        stage("Deploy 2 EC2 Instance"){
            steps{
                scripts{
                    ssh """
                    ssh -i ${SECOND_EC2_SSH_CREDENTIALS} ec2-user@ip << EOF 
                    docker pull myusername/myapp:${env.BUILD_ID}
                    docker stop myapp || true 
                    docker rm myapp || true 
                    docker run -d --name myapp -p 80:80 myusername/myap:${env.BUILD_ID}
                    EOF
                    """
                }
            }   
        }
        post { 
            always {
                cleanWs()
            }
        }
    }
}