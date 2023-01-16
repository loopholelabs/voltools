
cat fs_data_header > fs_data

for SIZE in `seq 10`
do
  truncate ${SIZE}g --size ${SIZE}g
  mkfs.ext4 ${SIZE}g
  go run . ${SIZE}g ${SIZE} >> fs_data
  rm ${SIZE}g
done

cat fs_data_footer >> fs_data
