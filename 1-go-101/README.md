# First GO Steps

## Go Modules

# Init a project
go mod init path_del_proyecto

# Verify the external code is not corrupt
go mod verify

# Replace source code
go mod edit -replace path_del_repo_online=path_del_repo_en_local

# Remove the replace
go mod edit -dropreplace path_del_repo_online

# Put together all the vendor code (third parties) used in our code
go mod vendor

# Remove all unused external packages
go mod tidy

# More about go modules
go help mod