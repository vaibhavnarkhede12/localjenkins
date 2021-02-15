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
      steps {
        sh 'go version'
      }
    }  
      
    stage('build') {
      steps {
        sh 'go build serve.go'
      }
    }
    
    stage('deploy') {
      steps {
        sh 'go build serve.go'
      }
    }
  }
}
