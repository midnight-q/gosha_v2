import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:gosha_app/common/validator.dart';
import 'package:gosha_app/controller/app_controller.dart';

class CreateAppScreen extends StatefulWidget {
  CreateAppScreen({Key? key}) : super(key: key);

  @override
  _CreateAppScreenState createState() => _CreateAppScreenState();
}

class _CreateAppScreenState extends State<CreateAppScreen> {
  final _formKey = GlobalKey<FormState>();

  final TextEditingController _email = TextEditingController();
  final TextEditingController _pass = TextEditingController();
  final TextEditingController _confirmPass = TextEditingController();

  AppController controller = Get.find();

  @override
  Widget build(BuildContext context) {
    return Center(
      child: SizedBox(
        width: 400,
        child: Form(
          key: _formKey,
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            crossAxisAlignment: CrossAxisAlignment.center,
            children: [
              Text(
                "Enter data for creating new app",
                style: TextStyle(fontSize: 18, fontWeight: FontWeight.w500),
              ),
              TextFormField(
                controller: _email,
                validator: validateEmail,
                decoration: const InputDecoration(
                  border: UnderlineInputBorder(),
                  labelText: 'Admin email',
                ),
              ),
              TextFormField(
                controller: _pass,
                obscureText: true,
                validator: validatePassword,
                decoration: const InputDecoration(
                  border: UnderlineInputBorder(),
                  labelText: 'Password',
                ),
              ),
              TextFormField(
                controller: _confirmPass,
                autovalidateMode: AutovalidateMode.onUserInteraction,
                obscureText: true,
                validator: (val) {
                  var res = validatePassword(val);
                  if (res != null) {
                    return res;
                  }
                  if (val != _pass.text) {
                    return "Password not math";
                  }
                  return null;
                },
                decoration: const InputDecoration(
                  border: UnderlineInputBorder(),
                  labelText: 'Repeat password',
                ),
              ),
              SizedBox(
                height: 50,
              ),
              SizedBox(
                width: 200,
                child: ElevatedButton(
                  onPressed: () {
                    if (!_formKey.currentState!.validate()) {
                      return;
                    }
                    controller.createApp(_email.text, _pass.text);
                  },
                  child: const Text(
                    'Create app',
                    style: TextStyle(fontSize: 16, fontWeight: FontWeight.w500),
                  ),
                ),
              )
            ],
          ),
        ),
      ),
    );
  }
}
