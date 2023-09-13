package service

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"os/exec"
)

var config map[string]interface{}

func Init() {
	// 转换为map
	data, _ := readYaml("config.yaml")
	err := yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("unmarshal to map failed: %v", err)
	}
}

func DoPull(action string) {
	// 去同步我的文件
	go func() {
		item := config[action].(map[interface{}]interface{})
		dir := item["dir"].(string)
		project := item["project"].(string)
		forceUpdate := item["force-update"].(bool)
		git := item["git"].(string)
		branch := item["branch"].(string)

		if forceUpdate == true {
			// 删除目录
			rm := exec.Command("rm", "-rf", dir+project)
			output, err2 := rm.CombinedOutput()
			log.Printf("delete result :\n%s\n", string(output))
			if err2 != nil {
				log.Printf("sync failed with %s\n", err2)
			}
			cloneGit(dir, git, branch)
		} else {
			if _, err := os.Stat(dir + project); os.IsNotExist(err) {
				log.Printf("目录不存在，新建项目 %s\n", dir+project)
				//fmt.Println("目录不存在")
				cloneGit(dir, git, branch)
			} else {
				log.Printf("刷新项目 %s\n", dir+project)
				refreshGit(dir, project)
			}
		}
	}()
}

func refreshGit(dir string, project string) {
	cmd := exec.Command("git", "pull")
	cmd.Dir = dir + project
	out, err := cmd.CombinedOutput()
	log.Printf("sync result :\n%s\n", string(out))
	if err != nil {
		log.Printf("sync failed with %s\n", err)
	}
}

func cloneGit(dir string, git string, branch string) {
	//  git clone -b gh-page https://github.com/xxx/xxx.git
	cmd := exec.Command("git", "clone", "-b", branch, git)
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	log.Printf("sync result :\n%s\n", string(out))
	if err != nil {
		log.Printf("sync failed with %s\n", err)
	}
}

func readYaml(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}
