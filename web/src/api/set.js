import request from '@/utils/request';

export const showSet = data => {
    return request({
        url: 'api/set/show',
        method: 'post',
        data
    });
};

export const delSet = data => {
    return request({
        url: 'api/set/del',
        method: 'post',
        data
    });
};

export const addSet = data => {
    return request({
        url: 'api/set/add',
        method: 'post',
        data
    });
};