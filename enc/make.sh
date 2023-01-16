
for SIZE in 100m 200m 300m 400m 500m 600m 700m 800m 900m 1g 2g 3g 4g 5g 6g 7g 8g 9g 10g 11g 12g 13g 14g 15g 16g 17g 18g 19g 20g 21g 22g 23g 24g 25g 26g 27g 28g 29g 30g 31g 32g 33g 34g 35g 36g 37g 38g 39g 40g 41g 42g 43g 44g 45g 46g 47g 48g 49g 50g 60g 70g 80g 90g 100g
do
  truncate $SIZE --size $SIZE
  mkfs.ext4 $SIZE
  go run . $SIZE >> fs_data
  rm $SIZE
done
