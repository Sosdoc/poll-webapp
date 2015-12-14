/// <reference path="../typings/angularjs/angular.d.ts" />

namespace Polls {
	export class CreatePollController {

		private scope: ng.IScope;

		static $inject: Array<string> = ['$scope'];

		constructor(private $scope: ng.IScope) {
			this.scope = $scope;
		}

	}
}