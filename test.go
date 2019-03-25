file, err := os.Open("/tmp/afile.txt")
if err != nil {
    panic(err)
}
file.WriteAt([]byte{0x1}, 1)
