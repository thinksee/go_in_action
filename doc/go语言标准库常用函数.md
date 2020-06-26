1. os常用函数
- 导出函数
```go
func Hostname()(name string, err error) // Hostname返回内核提供的主机名
func Environ() []string                 // Environ返回表示环境变量的格式为"key=value"的字符串的切片拷贝
func Getenv(key string) string          // Getenv检索并返回名为key的环境变量的值
func Getpid() int                       // Getpid返回调用者所在进程的进程ID
func Exit(code int)                     // Exit让当前程序以给出的状态码code退出。一般来说，状态码0表示成功，非0表示出错。程序会立刻终止，defer的函数不会被执行
func Stat(name string) (fi FileInfo, err error) // 获取文件信息
func Getwd() (dir string, err error)    // Getwd返回一个对应当前工作目录的根路径
func Mkdir(name string, perm FileMode) error // 使用指定的权限和名称创建一个目录
func MkdirAll(path string, perm FileMode) error // 使用指定的权限和名称创建一个目录，包括任何必要的上级目录，并返回nil，否则返回错误
func Remove(name string) error          // 删除name指定的文件或目录
func TempDir() string                   // 返回一个用于保管临时文件的默认目录
var Args []string                       // Args保管了命令行参数，第一个是程序名。
```
- File 结构体

```go
func Create(name string) (file *File, err error) // Create采用模式0666（任何人都可读写，不可执行）创建一个名为name的文件，如果文件已存在会截断它（为空文件）
func Open(name string) (file *File, err error) // Open打开一个文件用于读取。如果操作成功，返回的文件对象的方法可用于读取数据；对应的文件描述符具有O_RDONLY模式
func (f *File) Stat() (fi FileInfo, err error) // Stat返回描述文件f的FileInfo类型值
func (f *File) Readdir(n int) (fi []FileInfo, err error) // Readdir读取目录f的内容，返回一个有n个成员的[]FileInfo，这些FileInfo是被Lstat返回的，采用目录顺序
func (f *File) Read(b []byte) (n int, err error) // Read方法从f中读取最多len(b)字节数据并写入b
func (f *File) WriteString(s string) (ret int, err error) // 向文件中写入字符串
func (f *File) Sync() (err error) // Sync递交文件的当前内容进行稳定的存储。一般来说，这表示将文件系统的最近写入的数据在内存中的拷贝刷新到硬盘中稳定保存
func (f *File) Close() error // Close关闭文件f，使文件不能用于读写
```
- FileInfo 结构体
```go
type FileInfo interface {
    Name() string       // base name of the file
    Size() int64        // length in bytes for regular files; system-dependent for others
    Mode() FileMode     // file mode bits
    ModTime() time.Time // modification time
    IsDir() bool        // abbreviation for Mode().IsDir()
    Sys() interface{}   // underlying data source (can return nil)
}
```
2. Logger结构体
