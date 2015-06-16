'use strict';

angular.module('frontend')
  .controller('MainCtrl', function ($scope, Restangular) {
    var todos = Restangular.all("todos");
    todos.getList().then(onGetTodos);
    function onGetTodos(r) {
      $scope.todos = r;
    }
    $scope.handleSubmit = function() {
      $scope.newTodo.id = parseInt(_.uniqueId());
      todos.post($scope.newTodo);
      $scope.todos.push(angular.copy($scope.newTodo));
      $scope.newTodo = {};
    };
    $scope.onStatusChange = function(todo) {
      todo.put();
    };
  });
