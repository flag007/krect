# krect

```
wget https://raw.githubusercontent.com/arkadiyt/bounty-targets-data/master/data/domains.txt
mv domains.txt all.txt
go get -u github.com/flag007/krect

for i in $(cat all.txt); do waybackurls $i | grep '=http' | krect >> 1.txt; done
cat 1.txt | fff -s 302 -s 301 --output all
for i in $(find all -type f -name '*.headers'); do echo $i ; cat $i | grep "Location: https://google.com" ; done | grep "Location: https://google.com" -B 1 | tee out.log
```

