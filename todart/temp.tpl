import 'dart:convert';

class {{.ClassName}} {
  //Field Define
  {{call .genField}}

 //construct function define
  {{.ClassName}}({
    {{call .genConstruct}}
  });

  {{.ClassName}} copyWith({
    {{call .genCopyWithInputFun}}
  }) {
    return new {{.ClassName}}(
      {{call .genCopyWithRetFun}}
    );
  }

  String toJson() => json.encode(toMap());

  factory Dict.fromJson(String source) => Dict.fromMap(json.decode(source));

  @override
  String toString() {
    return '{{.ClassName}}{ {{- call .genToStringFun}}}';
  }

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      (other is {{.ClassName}} &&
          {{call .genOptEQFun}});

  @override
  int get hashCode =>
      {{call .genHashCodeFun}};

  factory {{.ClassName}}.fromMap(Map<String?, dynamic> map) {
    return new Article(
      {{call .genFromMapFun}}
    );
  }

  Map<String, dynamic> toMap() {
    // ignore: unnecessary_cast
    return {
      {{call .genToMapFun}}
    } as Map<String, dynamic>;
  }
//</editor-fold>
}