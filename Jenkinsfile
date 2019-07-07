pipeline {
    agent {
        docker { image 'golang:1.12.4' }
    }
    stages {
        stage('Test') {
            steps {
                sh 'go test ./...'
            }
        }
    }
}
