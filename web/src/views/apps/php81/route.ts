import type { RouteType } from '~/types/router'

const Layout = () => import('@/layout/IndexView.vue')

export default {
  name: 'php81',
  path: '/apps/php81',
  component: Layout,
  isHidden: true,
  children: [
    {
      name: 'apps-php81-index',
      path: '',
      component: () => import('../php/IndexView.vue'),
      meta: {
        title: 'PHP 8.1',
        icon: 'mdi:language-php',
        role: ['admin'],
        requireAuth: true,
        php: 81
      }
    }
  ]
} as RouteType