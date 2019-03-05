### grpc demo

## 参考资料：

[PHP implementation of gRPC](https://github.com/grpc/grpc/tree/master/src/php "PHP implementation of gRPC").

[Install golang](https://www.jianshu.com/p/85e98e9b003d "Install golang").

[grpc demo tutorial](https://www.jianshu.com/p/74b6080c3c5f "grpc demo tutorial").

    生成go代码：
    mkdir mr_gun
    protoc --go_out=plugins=grpc:./mr_gun mr_gun.proto

    生成php代码：
    mkdir phpClient
    protoc --proto_path=./ \
           --php_out=phpClient \
           --grpc_out=phpClient \
           --plugin=protoc-gen-grpc=./grpc/bins/opt/grpc_php_plugin \
           mr_gun.proto

## mysql数据库：

    create database mr_gun;

    create table gun(
      name varchar(20) not null primary key comment '名称',
      damage tinyint(1) unsigned not null comment '每发子弹伤害',
      bullets_shot tinyint(1) unsigned not null comment '每次射击子弹发数',
      total_damage tinyint(1) unsigned not null comment '每次射击总伤害',
      scatter tinyint(1) unsigned not null default 0 comment '是否散弹'
    ) engine Innodb default charset = utf8 comment 'MR.GUN';

    insert into gun(name, damage, bullets_shot, total_damage,scatter) values ('USP-S',3,1,3,0),
    ('DESERT EAGLE',8,1,8,0),
    ('.357 MAGNUM',6,1,6,0),
    ('AK-47',3,3,9,0),
    ('STEYR AUG',4,2,8,0),
    ('EVO 3',2,3,6,0),
    ('MODEL 1887',3,4,12,1),
    ('SCOUT',5,1,5,0),
    ('M240',3,4,12,0),
    ('NAVAL',7,1,7,0),
    ('VINTAGE',5,1,5,0),
    ('UZI',1,4,4,0),
    ('FAMAS',4,3,12,0),
    ('M16 CARBINE',3,3,9,0),
    ('KSG-12',5,3,15,1),
    ('HK556',6,2,12,0),
    ('BREN 2',2,5,10,0),
    ('SOCOM II',9,1,9,0),
    ('NIGHTHAWK',8,1,8,0),
    ('XLR-5',2,6,12,1),
    ('KELTEC M43',5,2,10,0),
    ('VAC',10,1,10,0),
    ('SCAR-L',4,3,12,0),
    ('A1-82',12,1,12,0),
    ('ARX-100',4,4,16,0),
    ('L96 AWP',10,1,10,0),
    ('SOLAR GUN',3,4,12,0),
    ('DOUBLE TROUBLE',7,2,14,0),
    ('CHARGER',9,1,9,0),
    ('PROBLEM SOLVER',4,4,16,0),
    ('GOODBYE',3,5,15,1),
    ('FOR FUN',5,3,15,0),
    ('SUPRISE',3,4,12,0),
    ('FIRESTARTER',8,1,8,0),
    ('WRATH',11,1,11,0),
    ('WHY NOT',4,3,12,1),
    ('ICE BREAKER',4,4,16,0),
    ('SHOW-OFF',6,3,18,1),
    ('BABY MONSTER',12,1,12,0),
    ('STEALTHY BOB',16,1,16,0),
    ('BIG MOUTH',6,4,24,0),
    ('THE UPGRADE',7,3,21,0);