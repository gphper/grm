import request from '@/utils/request';

export const showHash = data => {
    return request({
        url: 'grmapix/hash/show',
        method: 'post',
        data
    });
};

export const delHash = data => {
    return request({
        url: 'grmapix/hash/del',
        method: 'post',
        data
    });
};

export const addHash = data => {
    return request({
        url: 'grmapix/hash/add',
        method: 'post',
        data
    });
};