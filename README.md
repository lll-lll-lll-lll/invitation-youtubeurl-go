# Zenn
https://zenn.dev/jordan/articles/db2c4fd08e7387
# 経緯
料理本とかで時々「Youtubeにも解説記事載せています！」と書いている本があるが、QRコードが貼っていないのが意外と不便だと感じた。<br/>
情報関係に日頃から触れていない人からするとYoutubeURLを限定配信にし、そのURLをQRコードにするという発想は以外に思いつきにくいんじゃないかと思った<br/>
なので、IDとパスワードを入力すると招待コードが生成され、その招待コードとパスワード、IDを入力するとYoutubeのURLが生成されるアプリを作ってみたくなった<br/>

# 機能
- ユーザ作成機能
- 認証機能
- 招待コードの生成
- 招待コード生成時にaes暗号アルゴリズムを採用し、招待コードとidとパスワードが正しくないとyoutubeURLが返ってこない実装

# 手こずったところ
- aesで暗号化した際にbyteで返ってくるのだがそのbyteをstring()メソッドでutf8コード文字列に変えてもshiftjisにない文字が含まれている可能性があり、dbにインサートしようとすると文字化けが起きていてエラーが返ってくる<br/>
なのでhexパッケージを使って16進数にエンコードし、インサートするようにした。

### db接続
psql -h localhost -U app_user -d app_db

## 必要な設定
- Firebaseのシークレット情報を含んだjsonファイルをbackendディレクトリ内に入れる(account.json)


#　処理の流れ
### 1
**ユーザ作成**<br/>
<img width="778" alt="スクリーンショット 2022-12-09 0 28 59" src="https://user-images.githubusercontent.com/63499912/206486551-f99cbf6e-4f02-475b-b29d-0419f6566b26.png">

### 2
firebaseからfirebaseIDトークン取得し、そのトークンをAuthにセット<br/>
レスで招待コードを取得<br/>
<img width="800" alt="スクリーンショット 2022-12-09 0 36 05" src="https://user-images.githubusercontent.com/63499912/206488661-320597dc-0904-49a1-b69b-877c5f7a5398.png">


### 3
さっき入力したidパスワード、招待コードを入力。正しい場合YoutubeURLが取得できる

<img width="800" alt="スクリーンショット 2022-12-09 0 36 41" src="https://user-images.githubusercontent.com/63499912/206488844-bbe6f479-2d4a-4aa1-868d-7dd5c1844966.png">

