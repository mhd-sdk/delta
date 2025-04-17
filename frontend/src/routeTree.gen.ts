/* eslint-disable */

// @ts-nocheck

// noinspection JSUnusedGlobalSymbols

// This file was automatically generated by TanStack Router.
// You should NOT make any changes in this file as it will be overwritten.
// Additionally, you should also exclude this file from your linter and/or formatter to prevent it from being checked or modified.

// Import Routes

import { Route as rootRoute } from './routes/__root'
import { Route as AuthenticatedRouteImport } from './routes/_authenticated/route'
import { Route as AuthenticatedIndexImport } from './routes/_authenticated/index'
import { Route as errors503Import } from './routes/(errors)/503'
import { Route as errors500Import } from './routes/(errors)/500'
import { Route as errors404Import } from './routes/(errors)/404'
import { Route as errors403Import } from './routes/(errors)/403'
import { Route as errors401Import } from './routes/(errors)/401'
import { Route as authSignUpImport } from './routes/(auth)/sign-up'
import { Route as authSignIn2Import } from './routes/(auth)/sign-in-2'
import { Route as authSignInImport } from './routes/(auth)/sign-in'
import { Route as authOtpImport } from './routes/(auth)/otp'
import { Route as authForgotPasswordImport } from './routes/(auth)/forgot-password'
import { Route as AuthenticatedSettingsRouteImport } from './routes/_authenticated/settings/route'
import { Route as AuthenticatedUsersIndexImport } from './routes/_authenticated/users/index'
import { Route as AuthenticatedTasksIndexImport } from './routes/_authenticated/tasks/index'
import { Route as AuthenticatedSettingsIndexImport } from './routes/_authenticated/settings/index'
import { Route as AuthenticatedHelpCenterIndexImport } from './routes/_authenticated/help-center/index'
import { Route as AuthenticatedAlgorithmsIndexImport } from './routes/_authenticated/algorithms/index'
import { Route as AuthenticatedAlertsIndexImport } from './routes/_authenticated/alerts/index'
import { Route as AuthenticatedSettingsNotificationsImport } from './routes/_authenticated/settings/notifications'
import { Route as AuthenticatedSettingsDisplayImport } from './routes/_authenticated/settings/display'
import { Route as AuthenticatedSettingsAppearanceImport } from './routes/_authenticated/settings/appearance'
import { Route as AuthenticatedSettingsAccountImport } from './routes/_authenticated/settings/account'

// Create/Update Routes

const AuthenticatedRouteRoute = AuthenticatedRouteImport.update({
  id: '/_authenticated',
  getParentRoute: () => rootRoute,
} as any)

const AuthenticatedIndexRoute = AuthenticatedIndexImport.update({
  id: '/',
  path: '/',
  getParentRoute: () => AuthenticatedRouteRoute,
} as any)

const errors503Route = errors503Import.update({
  id: '/(errors)/503',
  path: '/503',
  getParentRoute: () => rootRoute,
} as any)

const errors500Route = errors500Import.update({
  id: '/(errors)/500',
  path: '/500',
  getParentRoute: () => rootRoute,
} as any)

const errors404Route = errors404Import.update({
  id: '/(errors)/404',
  path: '/404',
  getParentRoute: () => rootRoute,
} as any)

const errors403Route = errors403Import.update({
  id: '/(errors)/403',
  path: '/403',
  getParentRoute: () => rootRoute,
} as any)

const errors401Route = errors401Import.update({
  id: '/(errors)/401',
  path: '/401',
  getParentRoute: () => rootRoute,
} as any)

const authSignUpRoute = authSignUpImport.update({
  id: '/(auth)/sign-up',
  path: '/sign-up',
  getParentRoute: () => rootRoute,
} as any)

const authSignIn2Route = authSignIn2Import.update({
  id: '/(auth)/sign-in-2',
  path: '/sign-in-2',
  getParentRoute: () => rootRoute,
} as any)

const authSignInRoute = authSignInImport.update({
  id: '/(auth)/sign-in',
  path: '/sign-in',
  getParentRoute: () => rootRoute,
} as any)

const authOtpRoute = authOtpImport.update({
  id: '/(auth)/otp',
  path: '/otp',
  getParentRoute: () => rootRoute,
} as any)

const authForgotPasswordRoute = authForgotPasswordImport.update({
  id: '/(auth)/forgot-password',
  path: '/forgot-password',
  getParentRoute: () => rootRoute,
} as any)

const AuthenticatedSettingsRouteRoute = AuthenticatedSettingsRouteImport.update(
  {
    id: '/settings',
    path: '/settings',
    getParentRoute: () => AuthenticatedRouteRoute,
  } as any,
)

const AuthenticatedUsersIndexRoute = AuthenticatedUsersIndexImport.update({
  id: '/users/',
  path: '/users/',
  getParentRoute: () => AuthenticatedRouteRoute,
} as any)

const AuthenticatedTasksIndexRoute = AuthenticatedTasksIndexImport.update({
  id: '/tasks/',
  path: '/tasks/',
  getParentRoute: () => AuthenticatedRouteRoute,
} as any)

const AuthenticatedSettingsIndexRoute = AuthenticatedSettingsIndexImport.update(
  {
    id: '/',
    path: '/',
    getParentRoute: () => AuthenticatedSettingsRouteRoute,
  } as any,
)

const AuthenticatedHelpCenterIndexRoute =
  AuthenticatedHelpCenterIndexImport.update({
    id: '/help-center/',
    path: '/help-center/',
    getParentRoute: () => AuthenticatedRouteRoute,
  } as any)

const AuthenticatedAlgorithmsIndexRoute =
  AuthenticatedAlgorithmsIndexImport.update({
    id: '/algorithms/',
    path: '/algorithms/',
    getParentRoute: () => AuthenticatedRouteRoute,
  } as any)

const AuthenticatedAlertsIndexRoute = AuthenticatedAlertsIndexImport.update({
  id: '/alerts/',
  path: '/alerts/',
  getParentRoute: () => AuthenticatedRouteRoute,
} as any)

const AuthenticatedSettingsNotificationsRoute =
  AuthenticatedSettingsNotificationsImport.update({
    id: '/notifications',
    path: '/notifications',
    getParentRoute: () => AuthenticatedSettingsRouteRoute,
  } as any)

const AuthenticatedSettingsDisplayRoute =
  AuthenticatedSettingsDisplayImport.update({
    id: '/display',
    path: '/display',
    getParentRoute: () => AuthenticatedSettingsRouteRoute,
  } as any)

const AuthenticatedSettingsAppearanceRoute =
  AuthenticatedSettingsAppearanceImport.update({
    id: '/appearance',
    path: '/appearance',
    getParentRoute: () => AuthenticatedSettingsRouteRoute,
  } as any)

const AuthenticatedSettingsAccountRoute =
  AuthenticatedSettingsAccountImport.update({
    id: '/account',
    path: '/account',
    getParentRoute: () => AuthenticatedSettingsRouteRoute,
  } as any)

// Populate the FileRoutesByPath interface

declare module '@tanstack/react-router' {
  interface FileRoutesByPath {
    '/_authenticated': {
      id: '/_authenticated'
      path: ''
      fullPath: ''
      preLoaderRoute: typeof AuthenticatedRouteImport
      parentRoute: typeof rootRoute
    }
    '/_authenticated/settings': {
      id: '/_authenticated/settings'
      path: '/settings'
      fullPath: '/settings'
      preLoaderRoute: typeof AuthenticatedSettingsRouteImport
      parentRoute: typeof AuthenticatedRouteImport
    }
    '/(auth)/forgot-password': {
      id: '/(auth)/forgot-password'
      path: '/forgot-password'
      fullPath: '/forgot-password'
      preLoaderRoute: typeof authForgotPasswordImport
      parentRoute: typeof rootRoute
    }
    '/(auth)/otp': {
      id: '/(auth)/otp'
      path: '/otp'
      fullPath: '/otp'
      preLoaderRoute: typeof authOtpImport
      parentRoute: typeof rootRoute
    }
    '/(auth)/sign-in': {
      id: '/(auth)/sign-in'
      path: '/sign-in'
      fullPath: '/sign-in'
      preLoaderRoute: typeof authSignInImport
      parentRoute: typeof rootRoute
    }
    '/(auth)/sign-in-2': {
      id: '/(auth)/sign-in-2'
      path: '/sign-in-2'
      fullPath: '/sign-in-2'
      preLoaderRoute: typeof authSignIn2Import
      parentRoute: typeof rootRoute
    }
    '/(auth)/sign-up': {
      id: '/(auth)/sign-up'
      path: '/sign-up'
      fullPath: '/sign-up'
      preLoaderRoute: typeof authSignUpImport
      parentRoute: typeof rootRoute
    }
    '/(errors)/401': {
      id: '/(errors)/401'
      path: '/401'
      fullPath: '/401'
      preLoaderRoute: typeof errors401Import
      parentRoute: typeof rootRoute
    }
    '/(errors)/403': {
      id: '/(errors)/403'
      path: '/403'
      fullPath: '/403'
      preLoaderRoute: typeof errors403Import
      parentRoute: typeof rootRoute
    }
    '/(errors)/404': {
      id: '/(errors)/404'
      path: '/404'
      fullPath: '/404'
      preLoaderRoute: typeof errors404Import
      parentRoute: typeof rootRoute
    }
    '/(errors)/500': {
      id: '/(errors)/500'
      path: '/500'
      fullPath: '/500'
      preLoaderRoute: typeof errors500Import
      parentRoute: typeof rootRoute
    }
    '/(errors)/503': {
      id: '/(errors)/503'
      path: '/503'
      fullPath: '/503'
      preLoaderRoute: typeof errors503Import
      parentRoute: typeof rootRoute
    }
    '/_authenticated/': {
      id: '/_authenticated/'
      path: '/'
      fullPath: '/'
      preLoaderRoute: typeof AuthenticatedIndexImport
      parentRoute: typeof AuthenticatedRouteImport
    }
    '/_authenticated/settings/account': {
      id: '/_authenticated/settings/account'
      path: '/account'
      fullPath: '/settings/account'
      preLoaderRoute: typeof AuthenticatedSettingsAccountImport
      parentRoute: typeof AuthenticatedSettingsRouteImport
    }
    '/_authenticated/settings/appearance': {
      id: '/_authenticated/settings/appearance'
      path: '/appearance'
      fullPath: '/settings/appearance'
      preLoaderRoute: typeof AuthenticatedSettingsAppearanceImport
      parentRoute: typeof AuthenticatedSettingsRouteImport
    }
    '/_authenticated/settings/display': {
      id: '/_authenticated/settings/display'
      path: '/display'
      fullPath: '/settings/display'
      preLoaderRoute: typeof AuthenticatedSettingsDisplayImport
      parentRoute: typeof AuthenticatedSettingsRouteImport
    }
    '/_authenticated/settings/notifications': {
      id: '/_authenticated/settings/notifications'
      path: '/notifications'
      fullPath: '/settings/notifications'
      preLoaderRoute: typeof AuthenticatedSettingsNotificationsImport
      parentRoute: typeof AuthenticatedSettingsRouteImport
    }
    '/_authenticated/alerts/': {
      id: '/_authenticated/alerts/'
      path: '/alerts'
      fullPath: '/alerts'
      preLoaderRoute: typeof AuthenticatedAlertsIndexImport
      parentRoute: typeof AuthenticatedRouteImport
    }
    '/_authenticated/algorithms/': {
      id: '/_authenticated/algorithms/'
      path: '/algorithms'
      fullPath: '/algorithms'
      preLoaderRoute: typeof AuthenticatedAlgorithmsIndexImport
      parentRoute: typeof AuthenticatedRouteImport
    }
    '/_authenticated/help-center/': {
      id: '/_authenticated/help-center/'
      path: '/help-center'
      fullPath: '/help-center'
      preLoaderRoute: typeof AuthenticatedHelpCenterIndexImport
      parentRoute: typeof AuthenticatedRouteImport
    }
    '/_authenticated/settings/': {
      id: '/_authenticated/settings/'
      path: '/'
      fullPath: '/settings/'
      preLoaderRoute: typeof AuthenticatedSettingsIndexImport
      parentRoute: typeof AuthenticatedSettingsRouteImport
    }
    '/_authenticated/tasks/': {
      id: '/_authenticated/tasks/'
      path: '/tasks'
      fullPath: '/tasks'
      preLoaderRoute: typeof AuthenticatedTasksIndexImport
      parentRoute: typeof AuthenticatedRouteImport
    }
    '/_authenticated/users/': {
      id: '/_authenticated/users/'
      path: '/users'
      fullPath: '/users'
      preLoaderRoute: typeof AuthenticatedUsersIndexImport
      parentRoute: typeof AuthenticatedRouteImport
    }
  }
}

// Create and export the route tree

interface AuthenticatedSettingsRouteRouteChildren {
  AuthenticatedSettingsAccountRoute: typeof AuthenticatedSettingsAccountRoute
  AuthenticatedSettingsAppearanceRoute: typeof AuthenticatedSettingsAppearanceRoute
  AuthenticatedSettingsDisplayRoute: typeof AuthenticatedSettingsDisplayRoute
  AuthenticatedSettingsNotificationsRoute: typeof AuthenticatedSettingsNotificationsRoute
  AuthenticatedSettingsIndexRoute: typeof AuthenticatedSettingsIndexRoute
}

const AuthenticatedSettingsRouteRouteChildren: AuthenticatedSettingsRouteRouteChildren =
  {
    AuthenticatedSettingsAccountRoute: AuthenticatedSettingsAccountRoute,
    AuthenticatedSettingsAppearanceRoute: AuthenticatedSettingsAppearanceRoute,
    AuthenticatedSettingsDisplayRoute: AuthenticatedSettingsDisplayRoute,
    AuthenticatedSettingsNotificationsRoute:
      AuthenticatedSettingsNotificationsRoute,
    AuthenticatedSettingsIndexRoute: AuthenticatedSettingsIndexRoute,
  }

const AuthenticatedSettingsRouteRouteWithChildren =
  AuthenticatedSettingsRouteRoute._addFileChildren(
    AuthenticatedSettingsRouteRouteChildren,
  )

interface AuthenticatedRouteRouteChildren {
  AuthenticatedSettingsRouteRoute: typeof AuthenticatedSettingsRouteRouteWithChildren
  AuthenticatedIndexRoute: typeof AuthenticatedIndexRoute
  AuthenticatedAlertsIndexRoute: typeof AuthenticatedAlertsIndexRoute
  AuthenticatedAlgorithmsIndexRoute: typeof AuthenticatedAlgorithmsIndexRoute
  AuthenticatedHelpCenterIndexRoute: typeof AuthenticatedHelpCenterIndexRoute
  AuthenticatedTasksIndexRoute: typeof AuthenticatedTasksIndexRoute
  AuthenticatedUsersIndexRoute: typeof AuthenticatedUsersIndexRoute
}

const AuthenticatedRouteRouteChildren: AuthenticatedRouteRouteChildren = {
  AuthenticatedSettingsRouteRoute: AuthenticatedSettingsRouteRouteWithChildren,
  AuthenticatedIndexRoute: AuthenticatedIndexRoute,
  AuthenticatedAlertsIndexRoute: AuthenticatedAlertsIndexRoute,
  AuthenticatedAlgorithmsIndexRoute: AuthenticatedAlgorithmsIndexRoute,
  AuthenticatedHelpCenterIndexRoute: AuthenticatedHelpCenterIndexRoute,
  AuthenticatedTasksIndexRoute: AuthenticatedTasksIndexRoute,
  AuthenticatedUsersIndexRoute: AuthenticatedUsersIndexRoute,
}

const AuthenticatedRouteRouteWithChildren =
  AuthenticatedRouteRoute._addFileChildren(AuthenticatedRouteRouteChildren)

export interface FileRoutesByFullPath {
  '': typeof AuthenticatedRouteRouteWithChildren
  '/settings': typeof AuthenticatedSettingsRouteRouteWithChildren
  '/forgot-password': typeof authForgotPasswordRoute
  '/otp': typeof authOtpRoute
  '/sign-in': typeof authSignInRoute
  '/sign-in-2': typeof authSignIn2Route
  '/sign-up': typeof authSignUpRoute
  '/401': typeof errors401Route
  '/403': typeof errors403Route
  '/404': typeof errors404Route
  '/500': typeof errors500Route
  '/503': typeof errors503Route
  '/': typeof AuthenticatedIndexRoute
  '/settings/account': typeof AuthenticatedSettingsAccountRoute
  '/settings/appearance': typeof AuthenticatedSettingsAppearanceRoute
  '/settings/display': typeof AuthenticatedSettingsDisplayRoute
  '/settings/notifications': typeof AuthenticatedSettingsNotificationsRoute
  '/alerts': typeof AuthenticatedAlertsIndexRoute
  '/algorithms': typeof AuthenticatedAlgorithmsIndexRoute
  '/help-center': typeof AuthenticatedHelpCenterIndexRoute
  '/settings/': typeof AuthenticatedSettingsIndexRoute
  '/tasks': typeof AuthenticatedTasksIndexRoute
  '/users': typeof AuthenticatedUsersIndexRoute
}

export interface FileRoutesByTo {
  '/forgot-password': typeof authForgotPasswordRoute
  '/otp': typeof authOtpRoute
  '/sign-in': typeof authSignInRoute
  '/sign-in-2': typeof authSignIn2Route
  '/sign-up': typeof authSignUpRoute
  '/401': typeof errors401Route
  '/403': typeof errors403Route
  '/404': typeof errors404Route
  '/500': typeof errors500Route
  '/503': typeof errors503Route
  '/': typeof AuthenticatedIndexRoute
  '/settings/account': typeof AuthenticatedSettingsAccountRoute
  '/settings/appearance': typeof AuthenticatedSettingsAppearanceRoute
  '/settings/display': typeof AuthenticatedSettingsDisplayRoute
  '/settings/notifications': typeof AuthenticatedSettingsNotificationsRoute
  '/alerts': typeof AuthenticatedAlertsIndexRoute
  '/algorithms': typeof AuthenticatedAlgorithmsIndexRoute
  '/help-center': typeof AuthenticatedHelpCenterIndexRoute
  '/settings': typeof AuthenticatedSettingsIndexRoute
  '/tasks': typeof AuthenticatedTasksIndexRoute
  '/users': typeof AuthenticatedUsersIndexRoute
}

export interface FileRoutesById {
  __root__: typeof rootRoute
  '/_authenticated': typeof AuthenticatedRouteRouteWithChildren
  '/_authenticated/settings': typeof AuthenticatedSettingsRouteRouteWithChildren
  '/(auth)/forgot-password': typeof authForgotPasswordRoute
  '/(auth)/otp': typeof authOtpRoute
  '/(auth)/sign-in': typeof authSignInRoute
  '/(auth)/sign-in-2': typeof authSignIn2Route
  '/(auth)/sign-up': typeof authSignUpRoute
  '/(errors)/401': typeof errors401Route
  '/(errors)/403': typeof errors403Route
  '/(errors)/404': typeof errors404Route
  '/(errors)/500': typeof errors500Route
  '/(errors)/503': typeof errors503Route
  '/_authenticated/': typeof AuthenticatedIndexRoute
  '/_authenticated/settings/account': typeof AuthenticatedSettingsAccountRoute
  '/_authenticated/settings/appearance': typeof AuthenticatedSettingsAppearanceRoute
  '/_authenticated/settings/display': typeof AuthenticatedSettingsDisplayRoute
  '/_authenticated/settings/notifications': typeof AuthenticatedSettingsNotificationsRoute
  '/_authenticated/alerts/': typeof AuthenticatedAlertsIndexRoute
  '/_authenticated/algorithms/': typeof AuthenticatedAlgorithmsIndexRoute
  '/_authenticated/help-center/': typeof AuthenticatedHelpCenterIndexRoute
  '/_authenticated/settings/': typeof AuthenticatedSettingsIndexRoute
  '/_authenticated/tasks/': typeof AuthenticatedTasksIndexRoute
  '/_authenticated/users/': typeof AuthenticatedUsersIndexRoute
}

export interface FileRouteTypes {
  fileRoutesByFullPath: FileRoutesByFullPath
  fullPaths:
    | ''
    | '/settings'
    | '/forgot-password'
    | '/otp'
    | '/sign-in'
    | '/sign-in-2'
    | '/sign-up'
    | '/401'
    | '/403'
    | '/404'
    | '/500'
    | '/503'
    | '/'
    | '/settings/account'
    | '/settings/appearance'
    | '/settings/display'
    | '/settings/notifications'
    | '/alerts'
    | '/algorithms'
    | '/help-center'
    | '/settings/'
    | '/tasks'
    | '/users'
  fileRoutesByTo: FileRoutesByTo
  to:
    | '/forgot-password'
    | '/otp'
    | '/sign-in'
    | '/sign-in-2'
    | '/sign-up'
    | '/401'
    | '/403'
    | '/404'
    | '/500'
    | '/503'
    | '/'
    | '/settings/account'
    | '/settings/appearance'
    | '/settings/display'
    | '/settings/notifications'
    | '/alerts'
    | '/algorithms'
    | '/help-center'
    | '/settings'
    | '/tasks'
    | '/users'
  id:
    | '__root__'
    | '/_authenticated'
    | '/_authenticated/settings'
    | '/(auth)/forgot-password'
    | '/(auth)/otp'
    | '/(auth)/sign-in'
    | '/(auth)/sign-in-2'
    | '/(auth)/sign-up'
    | '/(errors)/401'
    | '/(errors)/403'
    | '/(errors)/404'
    | '/(errors)/500'
    | '/(errors)/503'
    | '/_authenticated/'
    | '/_authenticated/settings/account'
    | '/_authenticated/settings/appearance'
    | '/_authenticated/settings/display'
    | '/_authenticated/settings/notifications'
    | '/_authenticated/alerts/'
    | '/_authenticated/algorithms/'
    | '/_authenticated/help-center/'
    | '/_authenticated/settings/'
    | '/_authenticated/tasks/'
    | '/_authenticated/users/'
  fileRoutesById: FileRoutesById
}

export interface RootRouteChildren {
  AuthenticatedRouteRoute: typeof AuthenticatedRouteRouteWithChildren
  authForgotPasswordRoute: typeof authForgotPasswordRoute
  authOtpRoute: typeof authOtpRoute
  authSignInRoute: typeof authSignInRoute
  authSignIn2Route: typeof authSignIn2Route
  authSignUpRoute: typeof authSignUpRoute
  errors401Route: typeof errors401Route
  errors403Route: typeof errors403Route
  errors404Route: typeof errors404Route
  errors500Route: typeof errors500Route
  errors503Route: typeof errors503Route
}

const rootRouteChildren: RootRouteChildren = {
  AuthenticatedRouteRoute: AuthenticatedRouteRouteWithChildren,
  authForgotPasswordRoute: authForgotPasswordRoute,
  authOtpRoute: authOtpRoute,
  authSignInRoute: authSignInRoute,
  authSignIn2Route: authSignIn2Route,
  authSignUpRoute: authSignUpRoute,
  errors401Route: errors401Route,
  errors403Route: errors403Route,
  errors404Route: errors404Route,
  errors500Route: errors500Route,
  errors503Route: errors503Route,
}

export const routeTree = rootRoute
  ._addFileChildren(rootRouteChildren)
  ._addFileTypes<FileRouteTypes>()

/* ROUTE_MANIFEST_START
{
  "routes": {
    "__root__": {
      "filePath": "__root.tsx",
      "children": [
        "/_authenticated",
        "/(auth)/forgot-password",
        "/(auth)/otp",
        "/(auth)/sign-in",
        "/(auth)/sign-in-2",
        "/(auth)/sign-up",
        "/(errors)/401",
        "/(errors)/403",
        "/(errors)/404",
        "/(errors)/500",
        "/(errors)/503"
      ]
    },
    "/_authenticated": {
      "filePath": "_authenticated/route.tsx",
      "children": [
        "/_authenticated/settings",
        "/_authenticated/",
        "/_authenticated/alerts/",
        "/_authenticated/algorithms/",
        "/_authenticated/help-center/",
        "/_authenticated/tasks/",
        "/_authenticated/users/"
      ]
    },
    "/_authenticated/settings": {
      "filePath": "_authenticated/settings/route.tsx",
      "parent": "/_authenticated",
      "children": [
        "/_authenticated/settings/account",
        "/_authenticated/settings/appearance",
        "/_authenticated/settings/display",
        "/_authenticated/settings/notifications",
        "/_authenticated/settings/"
      ]
    },
    "/(auth)/forgot-password": {
      "filePath": "(auth)/forgot-password.tsx"
    },
    "/(auth)/otp": {
      "filePath": "(auth)/otp.tsx"
    },
    "/(auth)/sign-in": {
      "filePath": "(auth)/sign-in.tsx"
    },
    "/(auth)/sign-in-2": {
      "filePath": "(auth)/sign-in-2.tsx"
    },
    "/(auth)/sign-up": {
      "filePath": "(auth)/sign-up.tsx"
    },
    "/(errors)/401": {
      "filePath": "(errors)/401.tsx"
    },
    "/(errors)/403": {
      "filePath": "(errors)/403.tsx"
    },
    "/(errors)/404": {
      "filePath": "(errors)/404.tsx"
    },
    "/(errors)/500": {
      "filePath": "(errors)/500.tsx"
    },
    "/(errors)/503": {
      "filePath": "(errors)/503.tsx"
    },
    "/_authenticated/": {
      "filePath": "_authenticated/index.tsx",
      "parent": "/_authenticated"
    },
    "/_authenticated/settings/account": {
      "filePath": "_authenticated/settings/account.tsx",
      "parent": "/_authenticated/settings"
    },
    "/_authenticated/settings/appearance": {
      "filePath": "_authenticated/settings/appearance.tsx",
      "parent": "/_authenticated/settings"
    },
    "/_authenticated/settings/display": {
      "filePath": "_authenticated/settings/display.tsx",
      "parent": "/_authenticated/settings"
    },
    "/_authenticated/settings/notifications": {
      "filePath": "_authenticated/settings/notifications.tsx",
      "parent": "/_authenticated/settings"
    },
    "/_authenticated/alerts/": {
      "filePath": "_authenticated/alerts/index.tsx",
      "parent": "/_authenticated"
    },
    "/_authenticated/algorithms/": {
      "filePath": "_authenticated/algorithms/index.tsx",
      "parent": "/_authenticated"
    },
    "/_authenticated/help-center/": {
      "filePath": "_authenticated/help-center/index.tsx",
      "parent": "/_authenticated"
    },
    "/_authenticated/settings/": {
      "filePath": "_authenticated/settings/index.tsx",
      "parent": "/_authenticated/settings"
    },
    "/_authenticated/tasks/": {
      "filePath": "_authenticated/tasks/index.tsx",
      "parent": "/_authenticated"
    },
    "/_authenticated/users/": {
      "filePath": "_authenticated/users/index.tsx",
      "parent": "/_authenticated"
    }
  }
}
ROUTE_MANIFEST_END */
