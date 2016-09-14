echo "start synchronous mode (in miliseconds)" >> temp.txt 
for run in {1..20}
do
  go run src/client/main.go True >> temp.txt 
done

echo "start asynchronous mode (in miliseconds)" >> temp.txt 
for run in {1..20}
do
  go run src/client/main.go False >> temp.txt 
done
