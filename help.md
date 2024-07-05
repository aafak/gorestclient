PS C:\Users\aafakmoh\OneDrive - Hewlett Packard Enterprise\github-repos\personel\gorestclient> go mod init github.com/aafak/gorestclient
go: creating new go.mod: module github.com/aafak/gorestclient
go: to add module requirements and sums:
        go mod tidy
PS C:\Users\aafakmoh\OneDrive - Hewlett Packard Enterprise\github-repos\personel\gorestclient> go mod tidy
PS C:\Users\aafakmoh\OneDrive - Hewlett Packard Enterprise\github-repos\personel\gorestclient> 
PS C:\Users\aafakmoh\OneDrive - Hewlett Packard Enterprise\github-repos\personel\gorestclient> 


ASIAPACIFIC+aafakmoh@HPE-5CG3360726 MINGW64 ~/OneDrive - Hewlett Packard Enterprise/github-repos/personel/gorestclient (main)
$ git commit -a -m "Adding go rest client"
[main 7192a92] Adding go rest client
 5 files changed, 131 insertions(+)
 create mode 100644 go.mod
 create mode 100644 help.md
 create mode 100644 main.go
 create mode 100644 main_test.go

ASIAPACIFIC+aafakmoh@HPE-5CG3360726 MINGW64 ~/OneDrive - Hewlett Packard Enterprise/github-repos/personel/gorestclient (main)
$ git push origin main
Enumerating objects: 9, done.
Counting objects: 100% (9/9), done.
Delta compression using up to 12 threads
Compressing objects: 100% (6/6), done.
Writing objects: 100% (7/7), 1.64 KiB | 1.64 MiB/s, done.
Total 7 (delta 1), reused 0 (delta 0), pack-reused 0
remote: Resolving deltas: 100% (1/1), done.
To github.com:aafak/gorestclient.git
   5225c75..7192a92  main -> main

ASIAPACIFIC+aafakmoh@HPE-5CG3360726 MINGW64 ~/OneDrive - Hewlett Packard Enterprise/github-repos/personel/gorestclient (main)

ASIAPACIFIC+aafakmoh@HPE-5CG3360726 MINGW64 ~/OneDrive - Hewlett Packard Enterprise/github-repos/personel/gorestclient (main)
$ git tag v1.0.0

ASIAPACIFIC+aafakmoh@HPE-5CG3360726 MINGW64 ~/OneDrive - Hewlett Packard Enterprise/github-repos/personel/gorestclient (main)
$ git push origin v1.0.0
Total 0 (delta 0), reused 0 (delta 0), pack-reused 0
To github.com:aafak/gorestclient.git
 * [new tag]         v1.0.0 -> v1.0.0

ASIAPACIFIC+aafakmoh@HPE-5CG3360726 MINGW64 ~/OneDrive - Hewlett Packard Enterprise/github-repos/personel/gorestclient (main)


****************To update existing code, again create new tag and push

ASIAPACIFIC+aafakmoh@HPE-5CG3360726 MINGW64 ~/OneDrive - Hewlett Packard Enterprise/github-repos/personel/gorestclient (main)
$ git status
On branch main
Your branch is up to date with 'origin/main'.

Changes not staged for commit:
  (use "git add/rm <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        modified:   help.md
        deleted:    main.go
        deleted:    main_test.go

Untracked files:
  (use "git add <file>..." to include in what will be committed)
        restclient.go
        restclient_test.go

no changes added to commit (use "git add" and/or "git commit -a")

ASIAPACIFIC+aafakmoh@HPE-5CG3360726 MINGW64 ~/OneDrive - Hewlett Packard Enterprise/github-repos/personel/gorestclient (main)
$ git add .

ASIAPACIFIC+aafakmoh@HPE-5CG3360726 MINGW64 ~/OneDrive - Hewlett Packard Enterprise/github-repos/personel/gorestclient (main)
$ git commit -a -m "Updated the pcakage details"
[main ad139ed] Updated the pcakage details
 3 files changed, 50 insertions(+), 15 deletions(-)
 rename main.go => restclient.go (77%)
 rename main_test.go => restclient_test.go (97%)

ASIAPACIFIC+aafakmoh@HPE-5CG3360726 MINGW64 ~/OneDrive - Hewlett Packard Enterprise/github-repos/personel/gorestclient (main)
$ git push origin main
Enumerating objects: 7, done.
Counting objects: 100% (7/7), done.
Delta compression using up to 12 threads
Compressing objects: 100% (5/5), done.
Writing objects: 100% (5/5), 1.83 KiB | 1.83 MiB/s, done.
Total 5 (delta 0), reused 0 (delta 0), pack-reused 0
To github.com:aafak/gorestclient.git
   7192a92..ad139ed  main -> main

ASIAPACIFIC+aafakmoh@HPE-5CG3360726 MINGW64 ~/OneDrive - Hewlett Packard Enterprise/github-repos/personel/gorestclient (main)
$ git tag v1.0.1

ASIAPACIFIC+aafakmoh@HPE-5CG3360726 MINGW64 ~/OneDrive - Hewlett Packard Enterprise/github-repos/personel/gorestclient (main)
$ git push origin v1.0.1
Total 0 (delta 0), reused 0 (delta 0), pack-reused 0
To github.com:aafak/gorestclient.git
 * [new tag]         v1.0.1 -> v1.0.1

ASIAPACIFIC+aafakmoh@HPE-5CG3360726 MINGW64 ~/OneDrive - Hewlett Packard Enterprise/github-repos/personel/gorestclient (main)
$
