import request from '@/utils/request';

export const showList = data => {
    return request({
        url: 'api/list/show',
        method: 'post',
        data
    });
};

export const delList = data => {
    return request({
        url: 'api/list/del',
        method: 'post',
        data
    });
};

export const addList = data => {
    return request({
        url: 'api/list/add',
        method: 'post',
        data
    });
};