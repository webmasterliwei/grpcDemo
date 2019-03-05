<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Mr_gun;

/**
 */
class MrGunClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param \Mr_gun\Request $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function GetRank(\Mr_gun\Request $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/mr_gun.MrGun/GetRank',
        $argument,
        ['\Mr_gun\Reply', 'decode'],
        $metadata, $options);
    }

}
