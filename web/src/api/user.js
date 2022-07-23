import request from '@/utils/request';

export const login = data => {
    return request({
        url: 'api/user/login',
        method: 'post',
        data
    });
};