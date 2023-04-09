1- Using GOPRIVATE from go version 1.13- The new GOPRIVATE environment variable indicates module paths that are not publicly available. It serves as the default value for the lower-level GONOPROXY and GONOSUMDB variables, which provide finer-grained control over which modules are fetched via proxy and verified using the checksum database.
  
2- Add the following in system variable- ```export GOPRIVATE="gitlab.com/idmabar,bitbucket.org/idmabar,github.com/idmabar"```. *Change the export value based on your company/org name.  
  
