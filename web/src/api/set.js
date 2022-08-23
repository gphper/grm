import request from '@/utils/request';

export const showSet = data => {
    return request({
        url: 'grmapix/set/show',
        method: 'post',
        data
    });
};

export const delSet = data => {
    return request({
        url: 'grmapix/set/del',
        method: 'post',
        data
    });
};

export const addSet = data => {
    return request({
        url: 'grmapix/set/add',
        method: 'post',
        data
    });
};