import request from '@/utils/request';

export const openDb = data => {
    return request({
        url: 'grmapix/index/open',
        method: 'post',
        data
    });
};

export const getKeys = data => {
    return request({
        url: 'grmapix/index/getkeys',
        method: 'post',
        data
    });
};

export const getKeyType = data => {
    return request({
        url: 'grmapix/index/getkeytype',
        method: 'post',
        data
    });
};

export const delKeys = data => {
    return request({
        url: 'grmapix/index/delkey',
        method: 'post',
        data
    });
};

export const ttlKey = data => {
    return request({
        url: 'grmapix/index/ttlkey',
        method: 'post',
        data
    });
};

export const serInfo = data => {
    return request({
        url: 'grmapix/index/serinfo',
        method: 'post',
        data
    });
}

export const luaRun = data => {
    return request({
        url: 'grmapix/index/luarun',
        method: 'post',
        data
    });
}

export const setting = data => {
    return request({
        url: 'grmapix/index/setting',
        method: 'post',
        data
    });
}

export const getSetting = () => {
    return request({
        url: 'grmapix/index/setting',
        method: 'get'
    });
};