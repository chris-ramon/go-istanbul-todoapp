'use strict';

angular.module('frontend', ['restangular', 'ui.router', 'ui.bootstrap'])
  .config(function ($stateProvider, $urlRouterProvider, RestangularProvider) {
    RestangularProvider.setBaseUrl("http://localhost:8080");
    $stateProvider
      .state('home', {
        url: '/',
        templateUrl: 'app/main/main.html',
        controller: 'MainCtrl'
      });

    $urlRouterProvider.otherwise('/');
  })
;
