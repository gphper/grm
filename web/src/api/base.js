import request from '@/utils/request';

export const getDBs = query => {
    return request({
        url: 'api/getdb',
        method: 'get',
        params: query
    });
};