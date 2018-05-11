run
```
1.95 real         2.12 user         0.24 sys
time node run.js

#1.49 real         2.73 user         0.32 sys
time java -cp out/production/skii com.company.Main

#1.14 real         1.86 user         0.13 sys
time ./run
```


Scala:
```
sbt assembly
time java -jar target/scala-2.12/skii-assembly-0.1.0-SNAPSHOT.jar ./map.txt
```
