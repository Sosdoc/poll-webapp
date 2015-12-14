/// <reference path="typings/angularjs/angular.d.ts" />
/// <reference path="controllers/create.ts" />	
var app = angular.module("poll-webapp", []);

app.controller("CreatePollController", Polls.CreatePollController);