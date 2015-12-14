var Polls;
(function (Polls) {
    var Poll = (function () {
        function Poll(title, answers) {
            this.title = title;
            this.answers = answers;
        }
        return Poll;
    })();
    var Answer = (function () {
        function Answer(text, other) {
            this.text = text;
            this.other = other;
        }
        return Answer;
    })();
})(Polls || (Polls = {}));
var Polls;
(function (Polls) {
    var CreatePollController = (function () {
        function CreatePollController($scope) {
            this.$scope = $scope;
            this.scope = $scope;
        }
        CreatePollController.$inject = ['$scope'];
        return CreatePollController;
    })();
    Polls.CreatePollController = CreatePollController;
})(Polls || (Polls = {}));
var app = angular.module("poll-webapp", []);
app.controller("CreatePollController", Polls.CreatePollController);
