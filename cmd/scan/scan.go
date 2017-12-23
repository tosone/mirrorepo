package scan

import (
	"log"

	"github.com/spf13/viper"
)

func Initialize(scanDir []string) {
	log.Println(viper.Get("DatabaseEngine"))
	//var err error
	//for _, dir := range scanDir {
	//	err = filepath.Walk(dir, func(p string, info os.FileInfo, err error) error {
	//		if err != nil {
	//			log.Panicln(err)
	//		}
	//		if strings.Index(p, ".git") != -1 {
	//			return err
	//		}
	//		if info.IsDir() {
	//			if com.IsDir(path.Join(p, ".git")) {
	//				cmd := exec.Command("git", "status")
	//				cmd.Dir = p
	//				stdout, err := cmd.StdoutPipe()
	//				if err != nil {
	//					log.Fatal(err)
	//				}
	//				if err := cmd.Start(); err != nil {
	//					log.Fatal(err)
	//				}
	//				{
	//					b, _ := ioutil.ReadAll(stdout)
	//					log.Println(string(b))
	//				}
	//
	//				if err := cmd.Wait(); err != nil {
	//					log.Fatal(err)
	//				}
	//				log.Println(p)
	//			}
	//		}
	//		return err
	//	})
	//}
	//
	//if err != nil {
	//	log.Panicln(err)
	//}
}
