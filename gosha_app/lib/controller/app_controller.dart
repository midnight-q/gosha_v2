import 'package:file_picker/file_picker.dart';
import 'package:get/get_state_manager/src/simple/get_controllers.dart';
import 'package:get/get.dart';
import 'package:gosha_app/sdk/application.dart';

class AppController extends GetxController {
  var currentPath = "".obs;
  var isPathSelected = false.obs;
  var isAppExist = false.obs;
  var currentApp = Application.empty().obs;

  void getNewPath() async {
    String? path = await FilePicker.platform
        .getDirectoryPath(dialogTitle: "Select directory with project");
    if (path == null) {
      return;
    }

    this.currentPath.value = path;
    this.isPathSelected.value = true;

    var app = await getCurrentApp(path);

    this.setApp(app);
  }

  void setApp(Application app) {
    this.currentApp.value = app;
    this.isAppExist.value = app.isAppExist();
  }

  String getActionText() {
    if (!this.isPathSelected.value) {
      return "Select dir";
    }
    if (!this.isAppExist.value) {
      return "Empty dir";
    }
    return this.currentApp.value.name;
  }

  void createApp(String email, String password) async {
    if (this.currentApp.value.isAppExist()) {
      throw("App already exist");
    }
    var app = Application(
        name: "newApp", email: email, password: password, databaseType: 1);

    var res = await createApplication(app, this.currentPath.value);
    
    this.setApp(res.theModel);
  }
}
