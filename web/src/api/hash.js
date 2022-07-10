import request from '@/utils/request';

export const showHash = data => {
    return request({
        url: 'api/hash/show',
        method: 'post',
        data
    });
};

export const delHash = data => {
    return request({
        url: 'api/hash/del',
        method: 'post',
        data
    });
};

export const addHash = data => {
    return request({
        url: 'api/hash/add',
        method: 'post',
        data
    });
};