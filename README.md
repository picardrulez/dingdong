`[![dingdong Build and Test](https://github.com/picardrulez/dingdong/actions/workflows/go.yml/badge.svg)](https://github.com/picardrulez/dingdong/actions/workflows/go.yml)

# dingdong   
## USAGE:  
### Default Mode  
By default, dingdong checks a port, and responds with a 0 if the port has a listener, and a 1 if it does not.  
  
Example:  
`dingdong www.google.com:80`  
`0`   

`dingdong www.google.com:83732828384`  
`1`. 

### Human Output
the -h option forces dingdong to respond with either "listening" or "not listening". 

Example:  
`dingdong -h www.google.com:80 `  
`Listening`
