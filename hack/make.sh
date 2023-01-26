

for SIZE in `seq 10 10 250`
do
  truncate ${SIZE}g --size ${SIZE}g
  mkfs.ext4 ${SIZE}g
  cat fs_data_header > out/fs_data_${SIZE}.go
  go run . ${SIZE}g ${SIZE} >> out/fs_data_${SIZE}.go
  cat fs_data_footer >> out/fs_data_${SIZE}.go
  rm ${SIZE}g
done
