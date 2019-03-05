<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: mr_gun.proto

namespace Mr_gun;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;
use Google\Protobuf\Internal\GPBWrapperUtils;

/**
 * Generated from protobuf message <code>mr_gun.Reply</code>
 */
class Reply extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>repeated .mr_gun.GunInfo gunInfo = 1;</code>
     */
    private $gunInfo;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Mr_gun\GunInfo[]|\Google\Protobuf\Internal\RepeatedField $gunInfo
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\MrGun::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>repeated .mr_gun.GunInfo gunInfo = 1;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getGunInfo()
    {
        return $this->gunInfo;
    }

    /**
     * Generated from protobuf field <code>repeated .mr_gun.GunInfo gunInfo = 1;</code>
     * @param \Mr_gun\GunInfo[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setGunInfo($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Mr_gun\GunInfo::class);
        $this->gunInfo = $arr;

        return $this;
    }

}

