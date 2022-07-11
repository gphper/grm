import request from '@/utils/request';

export const showStream = data => {
    return request({
        url: 'api/stream/show',
        method: 'post',
        data
    });
};

export const delStream = data => {
    return request({
        url: 'api/stream/del',
        method: 'post',
        data
    });
};

export const addStream = data => {
    return request({
        url: 'api/stream/add',
        method: 'post',
        data
    });
};