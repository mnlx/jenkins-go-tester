pipeline {
    agent any

    triggers {
        cron {'*/1 * * * *'}
    }
    stages {
        stage('Build') {
            steps {
                sh '''
                    chmod +x build.sh
                    ./build.sh
                '''
            }
        }

        stage('Run') {
            steps {
                sh '''
                docker-compose up
                '''
            }
        }
    }
}