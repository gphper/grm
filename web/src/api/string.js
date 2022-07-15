import request from '@/utils/request';

export const showString = data => {
    return request({
        url: 'api/string/show',
        method: 'post',
        data
    });
};

export const addString = data => {
    return request({
        url: 'api/string/add',
        method: 'post',
        data
    });
};