pipeline {
  agent any
    
  tools {go "golang"}
    
  parameters{
    choice(
    name:'language',
      choices:[
        'python',
        'golang',
        'java'
      ],
      description:'select the deployment language')
    choice(
      name:'env',
      choices:[
        'dev',
        'prod'
      ],
     description:'chose your environment'
    )
  }
  stages {
        
   
    stage('testversion') {
      when {
        allOf {
          not{
            branch 'master'
          }
        }
      }
      steps {
        sh 'go version'
        script{
        }
      }
    }  
      
    stage('build') {
      steps {
        sh 'go build serve.go'
      }
    }
    
    stage('deploy') {
      steps {
        sh 'go run serve.go'
      }
    }
  }
}
