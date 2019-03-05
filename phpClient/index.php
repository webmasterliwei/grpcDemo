<html>
<head>
    <style>
        body {
            line-height: 1.5;
            font-family: -apple-system,".SFNSText-Regular","San Francisco","Roboto","Segoe UI","Helvetica Neue","Lucida Grande",Arial,sans-serif;
            -webkit-font-feature-settings: "kern" 1,"kern";
            -moz-font-feature-settings: "kern" 1,"kern";
            font-feature-settings: "kern" 1,"kern";
            -webkit-font-kerning: normal;
            font-kerning: normal;
            -webkit-font-smoothing: antialiased;
            -moz-osx-font-smoothing: grayscale;
            background-color: #fff;
            color: #222;
            margin: 0;
            padding: 0;
        }
        .container {
            width: 80%;
            margin: 0 auto;
            text-align: center;
        }
        .header {
            font-size: 13px;
            padding-top: 10px;
            display: inline-block;
        }
        .header a {
            text-decoration: none;
            color: #222;
            float: left;
            margin: 0 10px 10px 0;
            border: 1px solid #bdc1c4;
            padding: 5px 10px;
        }
        .header a:hover {
            background: #d7dadb;
        }
        table {
            width: 100%;
            padding: 0;
            margin-bottom: 1em;
            border-collapse: collapse;
            font-size: .75rem;
        }
        table tr {
            border-top: 1px solid #eee;
            background-color: #fff;
            margin: 0;
            padding: 0;
        }
        table th {
            font-weight: 700;
            border: 1px solid #bdc1c4;
            background: #d7dadb;
            margin: 0;
            padding: .5em;
        }
        table td {
            border: 1px solid #d7dadb;
            margin: 0;
            padding: .5em;
        }
    </style>
</head>
<body>
<div class="container">
    <div class="header">
        <a <?php if ($_GET['sort'] == 1 || !$_GET['sort']) {?> style="background: #d7dadb" <?php } ?> href="/?sort=1">所有枪每次射击总伤害排名</a>
        <a <?php if ($_GET['sort'] == 2) {?> style="background: #d7dadb" <?php } ?> href="/?sort=2">非散弹枪每次射击总伤害排名</a>
        <a <?php if ($_GET['sort'] == 3) {?> style="background: #d7dadb" <?php } ?> href="/?sort=3">散弹枪每次射击总伤害排名</a>
        <a <?php if ($_GET['sort'] == 4) {?> style="background: #d7dadb" <?php } ?> href="/?sort=4">非连射枪每次射击总伤害排名</a>
        <a <?php if ($_GET['sort'] == 5) {?> style="background: #d7dadb" <?php } ?> href="/?sort=5">每次射击子弹发数排名</a>
    </div>
    <table>
        <thead><tr><th>RANK</th><th>NAME</th><th>DAMAGE</th><th>BULLETS SHOT</th><th>TOTAL DAMAGE</th><th>SCATTER</th></tr></thead>
        <tbody>
        <?php
        /**
         * Created by PhpStorm.
         * User: lwei
         * Date: 19-2-18
         * Time: 下午4:08
         */
        require dirname(__FILE__).'/vendor/autoload.php';
        include_once dirname(__FILE__).'/Mr_gun/MrGunClient.php';
        include_once dirname(__FILE__).'/Mr_gun/Request.php';
        include_once dirname(__FILE__).'/Mr_gun/Reply.php';
        include_once dirname(__FILE__).'/Mr_gun/GunInfo.php';
        include_once dirname(__FILE__).'/GPBMetadata/MrGun.php';
        $sort = $_GET['sort'] ?? 1;
        try {
            $client = new \Mr_gun\MrGunClient('localhost:50051', [
                'credentials' => Grpc\ChannelCredentials::createInsecure(),
            ]);
            $request = new \Mr_gun\Request();
            $request->setSort($sort);
            /** @var \Mr_gun\Reply $reply */
            list($reply, $status) = $client->GetRank($request)->wait();
            if ($status->code != 0) {
                echo $status->details;
                $client->close();
                exit;
            }
            $rank = 1;
            $iterator = $reply->getGunInfo()->getIterator();//->serializeToJsonString();
            while ($iterator->valid()) {
                $current = $iterator->current();
                echo '<tr><td>' . $rank . '</td><td>' . $current->getName() . '</td><td>' . $current->getDamage() . '</td><td>' . $current->getBulletsShot() . '</td><td>' . $current->getTotalDamage() . '</td><td>' . $current->getScatter() . '</td></tr>';
                $iterator->next();
                $rank++;
            }
        } catch (\Throwable | \Error | \Exception $e) {
            echo  "error : " . $e->getMessage();
        }
        $client->close();
        ?>

        </tbody>
    </table>
</div>
</body>
</html>
