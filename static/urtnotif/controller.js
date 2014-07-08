/**
 * Author: Abu Ashraf Masnun
 * URL: http://masnun.me
 */

var app = angular.module('App', []);
app.config(function ($interpolateProvider) {
    $interpolateProvider.startSymbol('//');
    $interpolateProvider.endSymbol('//');
});


function UrbanTerrorController($scope, $http) {
    $scope.server = {};
    
    $scope.updateServer = function () {
        $http.get("/api/json")
            .success(function (response) {

                var current_players = [];
                for (x in $scope.server.players) {
                    current_players.push($scope.server.players[x].name);
                }

                $scope.server = response;

                $scope.server.configs.forEach(function(config) {
                    if(config.key == 'sv_hostname') {
                        $scope.server.name = config.value;
                    }

                    if(config.key == 'mapname') {
                        $scope.server.mapname = config.value;
                    }



                });

                var new_players = [];
                for (x in $scope.server.players) {
                    if (current_players.indexOf($scope.server.players[x].name) == -1) {
                        new_players.push($scope.server.players[x]);
                    }
                }

                if (new_players.length > 0) {
                    if (new_players.length > 1) {
                        $scope.sendDesktopNotification(null, new_players.length + " new players joined ")
                    } else {
                        $scope.sendDesktopNotification(null, new_players[0].name + " entered joined")
                    }
                }
                
            })
    }


    
    $scope.startMonitoring = function () {
        $scope.timer = setTimeout(function () {
            $scope.updateServer();
            $scope.startMonitoring();
        }, 10000);


        $("a#start_mon").hide();
        $("a#stop_mon").show();
    }

    $scope.stopMonitoring = function () {
        clearTimeout($scope.timer);
        $("a#stop_mon").hide();
        $("a#start_mon").show();
    }

    $scope.startDN = function () {
        $scope.enableDesktopNotifications();
    }

    $scope.stopDN = function () {
        $scope.sendDesktopNotification('Urban Terror Server Monitor', 'Desktop Notification has been disabled!');
        $scope.desktop_notification = false;
        $("a#stop_dn").hide();
        $("a#start_dn").show();
    }

    // Helpers

    $scope.enableDesktopNotifications = function() {
        if (Notification.permission === "granted") { // 0 is PERMISSION_ALLOWED
            $scope.desktop_notification = true;
            $("a#start_dn").hide();
            $("a#stop_dn").show();
            $scope.sendDesktopNotification('Urban Terror Server Monitor', 'Desktop Notification has been enabled!');
        } else {
            Notification.requestPermission(function(permission) {
                if (permission === "granted") {
                    $scope.sendDesktopNotification('Urban Terror Server Monitor', 'Desktop Notification has been enabled!');
                } else {
                    alert("You have to allow desktop notification!")
                }
            });
        }
    }

    $scope.sendDesktopNotification= function(title, message) {
        var title = title || "Urban Terror";
        if ($scope.desktop_notification) {
            var notif = new Notification(title, {"body": message});
            notif.onshow = function () {
                setTimeout(function () {
                    notif.close();
                }, 5000);
            }
        }
    }


    $scope.updateServer()
    $scope.startMonitoring();

   
}

