(function(){
	var App = Ember.Application.create();

	Ember.Handlebars.registerBoundHelper('currentDate', function() {
		var currentTime = new Date()
  		return currentTime.getFullYear();
	});
})();
