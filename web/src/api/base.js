import request from '@/utils/request';

export const getConnList = query => {
    return request({
        url: 'api/conn/list',
        method: 'get',
        params: query
    });
};

export const delConn = data => {
    return request({
        url: 'api/conn/del',
        method: 'post',
        data
    });
};

export const addConn = data => {
    return request({
        url: 'api/conn/add',
        method: 'post',
        data
    });
};

export const testConn = data => {
    return request({
        url: 'api/conn/test',
        method: 'post',
        data
    });
};