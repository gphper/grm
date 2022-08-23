import request from '@/utils/request';

export const getConnList = query => {
    return request({
        url: 'grmapix/conn/list',
        method: 'get',
        params: query
    });
};

export const delConn = data => {
    return request({
        url: 'grmapix/conn/del',
        method: 'post',
        data
    });
};

export const addConn = data => {
    return request({
        url: 'grmapix/conn/add',
        method: 'post',
        data
    });
};

export const testConn = data => {
    return request({
        url: 'grmapix/conn/test',
        method: 'post',
        data
    });
};