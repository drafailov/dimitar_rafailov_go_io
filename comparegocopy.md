
func Copy(dst Writer, src Reader) (written int64, err error)

The copy built-in function(io.Copy() ) copies elements from a source into a destination.
One issue with Copy() is that you cannot guarantee a maximum number of bytes. For example, you may want copy a log file up to its current file size. If the log continues to grow during your copy then you’ll end up with more bytes than expected. In this case is better to use the CopyN() function to specify an exact number of bytes to be written. Another issue with Copy() is that it requires an allocation for the 32KB buffer on every call. If you are performing a lot of copies then you can reuse your own buffer by using CopyBuffer() instead.


func CopyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error)

CopyBuffer is identical to Copy except that it provids buffer. It could return an error for a zero-sized buffer.


func CopyN(dst Writer, src Reader, n int64) (written int64, err error)

CopyN is identical to Copy except that it copies n bytes.
It returns the number of bytes copied or the earliest error encountered while copying.

Example: 
reader := strings.NewReader("Clouway")
written, err := io.CopyN(os.Stdout, reader, 4) // copy up to 4 bytes
fmt.Printf("\n%d, %v", written, err)

print result is : Clou


For copying large files (300 - 400 MB) is proper to use using CopyBuffer( using bigger buffer size) or CopyN(load in memory).
