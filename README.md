# Learning-Go

## Settings

### Go

1. Goをインストール

	<pre>
	brew install go --with-cc-all
	</pre>

2. 環境変数を設定する

	<pre>
	export GOROOT=/usr/local/opt/go/libexec
	export GOPATH="$HOME/go"
	export PATH="$GOROOT/bin:$GOPATH/bin:$PATH"
	</pre>
	
	
### Gox - Simple Go Cross Compilation

1. Goxをインストール

	```
	$ go get github.com/mitchellh/gox
	...
	$ gox -h
	...
	```

2. ツールチェーンをビルドする

	```
	$ gox -build-toolchain
	...
	```


### Debugging

1. GDBインストール

	<pre>
	brew install https://raw.github.com/Homebrew/homebrew-dupes/master/gdb.rb
	</pre>
	
2. gdb-certの証明書をキーチェンで発行

	refs. http://safx-dev.blogspot.jp/2014/04/macgo.html

3. `/System/Library/LaunchDaemons/com.apple.taskgated.plist`を書き換える

	_dbをインストールしただけでは動かなくて、taskgated(8)の設定も変更しなければいけない。これはTigerからSnow Leopardで使われていた、プロセスのアクセス制御方式を受け付けるように設定するため。_
	
	refs. http://qiita.com/ymotongpoo/items/81d3c945483cae734122

	```
		    <key>ProgramArguments</key>
		    <array>
		        <string>/usr/libexec/taskgated</string>
		-       <string>-s</string>
		+       <string>-sp</string>
		    </array>
		</dict>
		</plist>
	```

4. 署名する

	```
	$ codesign -s gdb-cert /usr/local/Cellar/gdb/7.9.1/bin/gdb
	$ sudo killall taskgated
	```

5. gdbにエイリアスを

	<pre>
	$ alias ggdb=/usr/local/Cellar/gdb/7.7/bin/gdb
	</pre>
	
6. 確認する

	<pre>
	$ touch main
	$ vi main.go
	</pre>
	
	サンプルとして、フィボナッチを求める。
	
	<pre>
	package main
	 
	import "fmt"
	 
	func fibonacci(n int) int {
	    if n == 0 {
	        return 0
	    }
	    if n == 1 {
	        return 1
	    }
	    return fibonacci(n-1) + fibonacci(n-2)
	}
	 
	func main() {
	    fmt.Println("fibonacci")
	    fmt.Println(fibonacci(20))
	}
	</pre>
	
	ビルドしてデバッグしてみよう。
	
	<pre>
	$ go build -gcflags "-N -l"
	$ ggdb ./debugTest -d .
	</pre>
	
	bで、ブレークポイントを設置することができる。
	あとは、rとかbtとかcとかよくわからないが、以下のように出れば成功。
	
	```
	(gdb) b main.fibonacci
	Breakpoint 1 at 0x2000: file /Users/safx/src/golang/src/main/foo.go, line 5.
	(gdb) r
	Starting program: /Users/safx/src/golang/main
	fibonacci
	[New Thread 0x160f of process 92283]
	[New Thread 0x1803 of process 92283]
	[New Thread 0x1903 of process 92283]
	 
	Breakpoint 1, main.fibonacci (n=20, ~anon1=1) at /Users/safx/src/golang/src/main/foo.go:5
	5   func fibonacci(n int) int {
	(gdb) bt
	#0  main.fibonacci (n=20, ~anon1=1) at /Users/safx/src/golang/src/main/foo.go:5
	#1  0x0000000000002167 in main.main () at /Users/safx/src/golang/src/main/foo.go:17
	(gdb) c
	Continuing.
	Breakpoint 1, main.fibonacci (n=19, ~anon1=10) at /Users/safx/src/golang/src/main/foo.go:5
	5   func fibonacci(n int) int {
	(gdb) clear
	Deleted breakpoint 1 
	(gdb) c
	Continuing.
	6765
	[Inferior 1 (process 92631) exited normally]
	```
	
## Build

#### Goで普通にビルドする

1. ビルドする

    以下を実行すると、binの配下にバイナリ実行形式が生成されます。
    
    <pre>
    GOOS=linux GOARCH=arm GOARM=5 GOBIN=../../bin go install
    </pre>

#### Goxでビルドする

1. ビルドする

	```
	$ gox
	```