import request from '@/utils/request';

export const showZset = data => {
    return request({
        url: 'grmapix/zset/show',
        method: 'post',
        data
    });
};

export const delZset = data => {
    return request({
        url: 'grmapix/zset/del',
        method: 'post',
        data
    });
};

export const addZset = data => {
    return request({
        url: 'grmapix/zset/add',
        method: 'post',
        data
    });
};