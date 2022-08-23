import request from '@/utils/request';

export const showList = data => {
    return request({
        url: 'grmapix/list/show',
        method: 'post',
        data
    });
};

export const delList = data => {
    return request({
        url: 'grmapix/list/del',
        method: 'post',
        data
    });
};

export const addList = data => {
    return request({
        url: 'grmapix/list/add',
        method: 'post',
        data
    });
};