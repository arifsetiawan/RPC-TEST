echo "start synchronous mode (in miliseconds)" >> temp.txt 
for run in {1..20}
do
  go run src/clientdb/main.go True >> temp.txt 
done

echo "start asynchronous mode (in miliseconds)" >> temp.txt 
for run in {1..20}
do
  go run src/clientdb/main.go False >> temp.txt 
done
