pipeline {
    agent {
        // Su dung golang:1.12.4 tich hop ca git de co the chay
        // git voi go module
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
